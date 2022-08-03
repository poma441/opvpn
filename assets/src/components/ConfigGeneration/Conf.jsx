import React, { useState } from 'react'
import axios from 'axios'
import Radio from '@mui/material/Radio';
import RadioGroup from '@mui/material/RadioGroup';
import FormControlLabel from '@mui/material/FormControlLabel';
import FormControl from '@mui/material/FormControl';
import FormLabel from '@mui/material/FormLabel';
import TextField from '@mui/material/TextField';
import { Box } from '@mui/system';
import MenuItem from '@mui/material/MenuItem';
import Select from '@mui/material/Select';
import InputLabel from '@mui/material/InputLabel';
import Button from '@mui/material/Button';

export const Conf = (props) => {
  const tunnel_lvl = [
    {val: 'tap', desc: "TAP: Level 2"},
    {val: 'tun', desc: "TUN: Level 3"}
  ]
  const tunLvlList = tunnel_lvl.map((lvl) =>
    <FormControlLabel value={lvl.val} control={<Radio />} label={lvl.desc} />
  )
  const [lvl, setLvl] = useState(tunnel_lvl[0].val)
  const handleSetLvl = (event) => {
    setLvl(event.target.value);
    setAdapterName(event.target.value + '0')
  };
  const protocols = [
    { val: 'udp', desc: "UDP" },
    { val: 'tcp', desc: "TCP" }
  ]
  const protoList = protocols.map((prot) =>
    <FormControlLabel value={prot.val} control={<Radio />} label={prot.desc} />
  )
  const [proto, setProto] = useState(protocols[0].val)
  const handleSetProto = (event) => {
    setProto(event.target.value);
  };
  const best_ciphers = [
    'AES-256-GCM',
    'AES-128-GCM',
    'CHACHA20-POLY1305',
    'AES-256-CBC',
    'AES-192-CBC',
    'AES-128-CBC',
  ]
  const cipherList = best_ciphers.map((ciph) =>
    <MenuItem value={ciph}>{ciph}</MenuItem>
  )
  const [cipher, setCipher] = useState(best_ciphers[0])
  const handleSetCipher = (event) => {
    setCipher(event.target.value);
  };
  const [serverIP, setServerIP] = useState('')
  const handleSetServerIP = (event) => {
    setServerIP(event.target.value);
  };
  const [port, setPort] = useState('1194')
  const handleSetPort = (event) => {
    setPort(event.target.value);
  };
  const [adapterName, setAdapterName] = useState(lvl + '0')
  const handleSetAdapterName = (event) => {
    setAdapterName(event.target.value);
  };
  const [addrPool, setAddrPool] = useState('')
  const handleSetAddrPool = (event) => {
    setAddrPool(event.target.value);
  };
  const [netmask, setNetmask] = useState('255.255.255.0')
  const handleSetNetmask = (event) => {
    setNetmask(event.target.value);
  };
  const [route, setRoute] = useState('')
  const handleSetRoute = (event) => {
    setRoute(event.target.value);
  };
  return (
    <Box 
      component="div" 
      sx={{
        '& .MuiTextField-root': { m: 1, width: '50ch' },
        display: 'flex',
        flexDirection: 'column',
      }}
    >
      {/* <h1>{props.clientCount}</h1> */}
      <TextField id="outlined-basic" label="Server IP" variant="outlined" value={serverIP} onChange={handleSetServerIP}/>
      <TextField id="outlined-basic" label="Port" variant="outlined" value={port} onChange={handleSetPort}/>
      <FormControl>
        <FormLabel id="demo-radio-buttons-group-label">Tunnel level</FormLabel>
        <RadioGroup
          aria-labelledby="demo-radio-buttons-group-label"
          defaultValue={lvl}
          name="radio-buttons-group"
          onChange={handleSetLvl}
          value={lvl}
        >
          {tunLvlList}
        </RadioGroup>
        <FormLabel id="demo-radio-buttons-group-label">Protocol</FormLabel>
        <RadioGroup
          aria-labelledby="demo-radio-buttons-group-label"
          defaultValue={proto}
          name="radio-buttons-group"
          onChange={handleSetProto}
          value={proto}
        >
          {protoList}
        </RadioGroup>
      </FormControl>
      <TextField id="outlined-basic" label="Virtual adapter name" variant="outlined" value={adapterName} onChange={handleSetAdapterName}/>  
      <InputLabel id="demo-simple-select-label">Cipher</InputLabel>
      <Select
        labelId="demo-simple-select-label"
        id="demo-simple-select"
        value={cipher}
        label="Cipher"
        onChange={handleSetCipher}
        sx={{width: 300}}
      >
        {cipherList}
      </Select>
      <TextField id="outlined-basic" label="IP Address pool cherez zpt" variant="outlined" value={addrPool} onChange={handleSetAddrPool}/>
      <TextField id="outlined-basic" label="Netmask, example: 255.255.255.0" variant="outlined" value={netmask} onChange={handleSetNetmask}/>
      <TextField 
        sx={{width: 600}} 
        id="outlined-basic" 
        label="Route, example: route 10.1.0.0 255.255.0.0 10.1.254.1" 
        variant="outlined" 
        value={route}
        onChange={handleSetRoute}
      />
      <Button 
        variant='contained' 
        sx={{width: 200}} 
        onClick={async () => { 
          const response = await axios.post('http://localhost:8080/conf', {
            "serverIP": serverIP,
            "port": port,
            "proto": proto,
            "tunnel_lvl": lvl,
            "dev": adapterName,
            "cipher": cipher,
            "ifconfig-pool": addrPool,
            "netmask": netmask,
            "push": route 
          })
          // console.log(response.data)
          props.showLogs(response.data.msg)
        }}
      >Generate</Button>
    </Box>
  )
}


