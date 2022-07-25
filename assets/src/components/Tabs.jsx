import React from "react";
import { TabContent } from "./TabContent";
import '../styles/Tabs.css'


export default function Tabs({ items }) {
    const [ active, setActive ] = React.useState(0);
  
    const openTab = e => setActive(+e.target.dataset.index);
  
    return (
      <div>
        <div className="tab">
          {items.map((n, i) => (
            <button
              className={`tablinks ${i === active ? 'active' : ''}`}
              onClick={openTab}
              data-index={i}
            >{n.title}</button>
          ))}
        </div>
        {items[active] && <TabContent {...items[active]} />}
      </div>
    );
  }