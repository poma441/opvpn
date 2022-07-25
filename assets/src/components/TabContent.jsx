import React from 'react'
import { Conf } from './Conf'
import { Keys } from './Keys'
import { ServManage } from './ServManage'



export const TabContent = ({title, content}) => {
  return (
    <div>
    { ( title === "Keys" &&
         <Keys />) ||
      ( title === "Conf" &&
         <Conf/>) ||
      ( title === "Serv" &&
        <ServManage/>)
    }
      <h3>Logging</h3>
      
    
    </div>
  )
}
