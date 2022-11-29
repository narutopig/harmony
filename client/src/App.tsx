import React from "react";
import logo from "./logo.svg";
import "./App.css";
import { Link, Route, Routes } from "react-router-dom";
import Home from "./pages/Home";
import Another from "./pages/Another";

function App() {
  // the header tag can be extracted into a navbar for better layout
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <div className="menu">
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/another">Another Page</Link>
            </li>
          </ul>
        </div>
      </header>
      <div className="content">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/another" element={<Another />} />
        </Routes>
      </div>
    </div>
  );
}

export default App;
