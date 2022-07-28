import * as React from 'react';
import Button from '@mui/material/Button';
import ButtonGroup from '@mui/material/ButtonGroup';
import axios from 'axios'

export default function BasicButtonGroup() {
    async function sendStart () {
        const response = await axios.post('http://localhost:8080/management', {"cmd": "start"})
        console.log(response.data.msg)
    }
    async function sendStop () {
        const response = await axios.post('http://localhost:8080/management', {"cmd": "stop"})
        console.log(response.data.msg)
    }
    async function sendStatus () {
        const response = await axios.post('http://localhost:8080/management', {"cmd": "status"})
        console.log(response.data.msg)
    }

  return (
    <ButtonGroup variant="contained" aria-label="outlined primary button group">
      <Button onClick={sendStart}>Start</Button>
      <Button onClick={sendStop}>Stop</Button>
      <Button onClick={sendStatus}>Status</Button>
    </ButtonGroup>
  );
}
