package manage_server

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

// Инкапсулируем функцию для выполнения команд
func ExecCommand(commandName string, params []string) string {
	logging := ""

	// Выполнение заказа
	cmd := exec.Command(commandName, params...)

	// Отображение запущенных команд
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()
	errReader, errr := cmd.StderrPipe()

	if errr != nil {
		// fmt.Println("err:" + errr.Error())
		logging += "err: " + errr.Error()
	}

	// Включаем обработку ошибок
	go handlerErr(errReader)

	if err != nil {
		// fmt.Println(err)
		logging += err.Error()
		return logging
	}

	cmd.Start()
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		cmdRe := ConvertByte2String(in.Bytes(), "GB18030")
		logging += cmdRe
		// fmt.Println(cmdRe)
	}
	cmd.Wait()
	cmd.Wait()
	fmt.Println(logging)
	return logging
}

// Открываем сопрограмму с ошибкой
func handlerErr(errReader io.ReadCloser) {
	in := bufio.NewScanner(errReader)
	for in.Scan() {
		cmdRe := ConvertByte2String(in.Bytes(), "GB18030")
		fmt.Errorf(cmdRe)
	}
}

// Перекодировать символы
func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	// case GB18030:
	// 	var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
	// 	str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
