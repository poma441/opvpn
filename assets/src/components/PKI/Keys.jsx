import React from 'react'
// import { useState } from 'react';
import { useRef } from 'react'
import axios from 'axios'
import FormGroup from '@mui/material/FormGroup';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';
import Typography from '@mui/material/Typography';
import Slider from '@mui/material/Slider';
import MuiInput from '@mui/material/Input';
import { styled } from '@mui/material/styles';

const Input = styled(MuiInput)`
  width: 42px;
`;

export const Keys = () => {
    const inpRefCA = useRef();
    const inpRefSrv = useRef();

    const [clientCount, setClientCount] = React.useState(6);

    const handleSliderChange = (event, newClientCount) => {
      setClientCount(newClientCount);
    };

    const handleInputChange = (event) => {
      setClientCount(event.target.value === '' ? '' : Number(event.target.value));
    };

    const handleBlur = () => {
      if (clientCount < 0) {
        setClientCount(0);
      } else if (clientCount > 24) {
        setClientCount(24);
      }
    };

    async function sendInfo () {
        console.log(inpRefCA.current.checked)
        console.log(inpRefSrv.current.checked)
        console.log(clientCount)
        const response = await axios.post('http://localhost:8080/keys', {"ca": inpRefCA.current.checked,"server": inpRefSrv.current.checked,"clients": clientCount})
        console.log(response.data)
    }

  return (
    <div>
      <FormGroup>
        <FormControlLabel control={<Checkbox  inputRef={inpRefCA}/>} label="CA" />
        <FormControlLabel control={<Checkbox  inputRef={inpRefSrv}/>} label="Server" />
      </FormGroup>
      <Box sx={{ width: 250 }}>
        <Typography id="input-slider" gutterBottom>
          Clients:
        </Typography>
        <Grid container spacing={2} alignItems="center">
          <Grid item>
          </Grid>
          <Grid item xs>
            <Slider
              value={typeof clientCount === 'number' ? clientCount : 0}
              onChange={handleSliderChange}
              aria-labelledby="input-slider"
            />
          </Grid>
          <Grid item>
            <Input
              value={clientCount}
              size="small"
              onChange={handleInputChange}
              onBlur={handleBlur}
              inputProps={{
                step: 2,
                min: 0,
                max: 24,
                type: 'number',
                'aria-labelledby': 'input-slider',
              }}
            />
          </Grid>
        </Grid>
      </Box>
      <Button variant='contained' onClick={sendInfo}>Generate</Button>
    </div>
  )
}
