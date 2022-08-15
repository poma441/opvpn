#!/bin/bash
client_count=$1
 
#Проверка наличия директории openvpn если есть то удаляем и Создание заново, иначе Создание
if [[ -e /etc/openvpn ]]; then
  rm -rf /etc/openvpn
  mkdir /etc/openvpn
  mkdir /etc/openvpn/certs
  echo "Удалена старая директория openvpn, создана новая"
else
  mkdir /etc/openvpn
  mkdir /etc/openvpn/certs
  echo "Cоздана новая директория openvpn"
fi
#Copying easy-rsa to openvpn directory
sudo cp -R /usr/share/easy-rsa /etc/openvpn
 
#Create file named vars with user's configurations
touch /etc/openvpn/easy-rsa/vars
 
# Default vars
echo "Enter main configurations for certs creation"
echo "All values are not important besides certs validity"
 

country="RU"
key_size=2048
region="Moscow"
city="Moscow"
mail="example@example.com"
expire=3650
 
#Entering vars
cat <<EOF >/etc/openvpn/easy-rsa/vars
#### FOR Domain
####set_var EASYRSA_DN $domain_name
 
set_var EASYRSA_REQ_COUNTRY $country
set_var EASYRSA_KEY_SIZE $key_size
set_var EASYRSA_REQ_region $region
set_var EASYRSA_REQ_CITY $city
set_var EASYRSA_REQ_ORG $domain_name
set_var EASYRSA_REQ_EMAIL $mail
set_var EASYRSA_REQ_OU $domain_name
set_var EASYRSA_REQ_CN testCN
set_var EASYRSA_CERT_EXPIRE $expire
set_var EASYRSA_DH_KEY_SIZE $key_size
set_var EASYRSA_BATCH       "yes"
EOF
 
#Init PKI
cd /etc/openvpn/easy-rsa/ || exit; /etc/openvpn/easy-rsa/easyrsa init-pki
sudo dd if=/dev/urandom of=pki/.rand bs=256 count=1
sudo dd if=/dev/urandom of=pki/.rnd bs=256 count=1
#Создание ключа центра сертификации
/etc/openvpn/easy-rsa/easyrsa build-ca nopass
 
#Создание сертификата сервера
/etc/openvpn/easy-rsa/easyrsa build-server-full server nopass
 
#Создание файл Диффи Хелмана
/etc/openvpn/easy-rsa/easyrsa gen-dh
 
#Crl для информации об активных/отозванных сертификатов
/etc/openvpn/easy-rsa/easyrsa gen-crl
 
#Добавление HMAC
sudo openvpn --genkey secret /etc/openvpn/certs/ta.key
 
#Теперь копируем все что создали в папку certs
cp /etc/openvpn/easy-rsa/pki/ca.crt /etc/openvpn/easy-rsa/pki/crl.pem /etc/openvpn/easy-rsa/pki/dh.pem /etc/openvpn/certs/
cp /etc/openvpn/easy-rsa/pki/issued/server.crt /etc/openvpn/certs/
cp /etc/openvpn/easy-rsa/pki/private/server.key /etc/openvpn/certs/
cp /etc/openvpn/easy-rsa/pki/private/ca.key /etc/openvpn/certs/

create_client() {
  mkdir /etc/openvpn/certs/client$client_count
  cd /etc/openvpn/easy-rsa/
  /etc/openvpn/easy-rsa/easyrsa build-client-full "client$client_count" nopass
     cp /etc/openvpn/easy-rsa/pki/issued/client$client_count.crt /etc/openvpn/certs/client$client_count
     cp /etc/openvpn/easy-rsa/pki/private/client$client_count.key /etc/openvpn/certs/client$client_count
     cp /etc/openvpn/certs/ca.crt /etc/openvpn/certs/client$client_count
     cp /etc/openvpn/certs/dh.pem /etc/openvpn/certs/client$client_count
     cp /etc/openvpn/certs/ta.key /etc/openvpn/certs/client$client_count
} #Запуск функции создания клиентов по счетчику
while [[ $client_count -ne -1 ]]; do
  create_client
  let "client_count=$client_count-1"
done
