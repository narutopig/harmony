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
