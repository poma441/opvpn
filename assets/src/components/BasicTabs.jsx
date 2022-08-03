import * as React from 'react';
import TabPanel from './TabPanel';
import { Keys } from './PKI/Keys';
import { Conf } from './ConfigGeneration/Conf';
import BasicButtonGroup from './ServerManagement/BasicButtonGroup';
import PropTypes from 'prop-types';
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';
import { styled } from '@mui/material/styles';
import Paper from '@mui/material/Paper';
import axios from 'axios';



TabPanel.propTypes = {
  children: PropTypes.node,
  index: PropTypes.number.isRequired,
  value: PropTypes.number.isRequired,
};

function a11yProps(index) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}

export default function BasicTabs(props) {
  const StyledPaper = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(2),
    // maxWidth: 400,
    color: theme.palette.text.primary,
  }));
  const [value, setValue] = React.useState(0);

  const handleChange = (event, newValue) => {
    setValue(newValue);
  };

  const [logs, setLogs] = React.useState([])
  const addLogs = (newLogs) => {
    setLogs([newLogs, ...logs])
  }
  const showLogs = logs.map((data) =>
    <StyledPaper
      sx={{
        my: 1,
        p: 2,
      }}
    >
      <Grid container wrap="nowrap" spacing={2}>
        <Grid item xs>
          <Typography>{data}</Typography>
        </Grid>
      </Grid>
    </StyledPaper>
  )
  const [clientCount, setClientCount] = React.useState(0)
  const updateClientCount = (count) => {
    setClientCount(count)
  }
  const showDownloadLinks = (clientCount) => {
    let links = []
    for (let i = 0; i < clientCount; i++) {
        links.push(
            <button key={i} onClick={async () => {
                const response = await axios({
                  url: "http://localhost:8080/download?name=buffer.txt",
                  method: 'GET',
                  responseType: 'blob', // important
                })
                const url = window.URL.createObjectURL(new Blob([response.data]));
                const link = document.createElement('a');
                link.href = url;
                link.setAttribute('download', 'file.txt'); //or any other extension
                document.body.appendChild(link);
                link.click();
              } 
            }>Download config for client {i + 1}</button>
        )
    }
    return links
  }
    
  return (
    <Box sx={{ width: '100%' }}>
      <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
        <Tabs value={value} onChange={handleChange} aria-label="basic tabs example">
          <Tab label="Keys" {...a11yProps(0)} />
          <Tab label="Config" {...a11yProps(1)} />
          <Tab label="Server management" {...a11yProps(2)} />
        </Tabs>
      </Box>
      <TabPanel value={value} index={0}>
        <Keys showLogs={addLogs} updateClientCount={updateClientCount}/>
      </TabPanel>
      <TabPanel value={value} index={1}>
        <Conf showLogs={addLogs} clientCount={clientCount}/>
      </TabPanel>
      <TabPanel value={value} index={2}>
        <BasicButtonGroup showLogs={addLogs}/>
      </TabPanel>
      {showDownloadLinks(clientCount)}
      <h1>Logs</h1>
      {showLogs}
    </Box>
  );
}
