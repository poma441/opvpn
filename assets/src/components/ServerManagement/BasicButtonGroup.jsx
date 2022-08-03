import * as React from 'react';
import { useState } from 'react';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import ButtonGroup from '@mui/material/ButtonGroup';
// import axios from 'axios'
import Chip from '@mui/material/Chip';

export default function BasicButtonGroup(props) {

    const [serverStatus, setServerStatus] = useState('Stopped')
  //   async function sendCMD (cmd) {
  //     const response = await axios.post('http://localhost:8080/management', {"cmd": cmd})
  //     props.showLogs(response.data.msg)
  // }
  function sendCMD (cmd) {
    props.showLogs(cmd)
  }

  return (
    <Box>
      <ButtonGroup variant="contained" aria-label="outlined primary button group">
        <Button onClick={() => {
          sendCMD("start") 
          setServerStatus("Started")
          }}>Start</Button>
        <Button onClick={() => {sendCMD("stop"); setServerStatus("Stopped")}}>Stop</Button>
        <Button onClick={() => {sendCMD("status")}}>Status</Button>
      </ButtonGroup>
      <Chip label={serverStatus} color={serverStatus === "Started" ? "success" : "error"} />
    </Box>

  );
}
