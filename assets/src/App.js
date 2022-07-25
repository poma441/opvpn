import React from "react";
import Tabs from "./components/Tabs"

function App() {

  const items = [
    { title: 'Keys'},
    { title: 'Conf'},
    { title: 'Serv'},
  ];

  return (
    <div className="App">
      <Tabs items={items}/>
    </div>
  );
}

export default App;
