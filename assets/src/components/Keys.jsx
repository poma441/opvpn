import React from 'react'
import { useState } from 'react';
import { useRef } from 'react'
import axios from 'axios'

export const Keys = () => {
    const inpRefCA = useRef();
    const inpRefSrv = useRef();
    const [clientCount, setClientCount] = useState(6)

    async function sendInfo (e) {
        e.preventDefault()
        console.log(inpRefCA.current.checked)
        console.log(inpRefSrv.current.checked)
        console.log(clientCount)
        // const response = await axios.get('http://localhost:8080/keys')
        const response = await axios.post('http://localhost:8080/keys', {"id": "1","CA": true,"Server": true,"ClientCount": 6})
        console.log(response.data)
    }

  return (
    <form>
          <p>
            Certificate Authority (CA): 
            <input 
                
                type="checkbox" 
                ref={inpRefCA}
            />
          </p>
          <p>Server Certificate: 
            <input 
                type="checkbox"
                ref={inpRefSrv} 
            />
            </p>
          <p>Client count <input type="number" value={clientCount} onChange={e => setClientCount(e.target.value)}/></p>
          <button onClick={sendInfo}>Generate</button>
    </form>
  )
}
