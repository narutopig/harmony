import { useEffect } from "react";

import "./home.css";

function Home() {
  const channels = [
    "asdf",
    "asdf",
    "asdf",
    "asdf",
    "asdf",
    "asdf",
    "asdf",
    "asdf",
    "asdf",
    "asdf",
    "asdf",
  ];


  useEffect(() => {
    const ws = new WebSocket("wss://ws.bitstamp.net");

    ws.onopen = () => {
        console.log("connection")
    }

    ws.onmessage = (message) => {
        console.log(message)
    }

    return () => {
        ws.close(0, "User disconnected")
    }
  })

  return (
    <div className="container">
      <div id="channels" className="content">
        {channels.map((channel) => (
          <h1>#{channel}</h1>
        ))}
      </div>
      <div id="messages" className="content">
        <h1>Hello World!</h1>
      </div>
    </div>
  );
}

export default Home;
