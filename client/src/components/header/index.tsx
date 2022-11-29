import { Link } from "react-router-dom";
import "./header.css";

function Header() {
  return (
    <div className="App-header">
      <div className="menu">
        <h1>
          <Link to="/">Harmony</Link>
        </h1>
      </div>
    </div>
  );
}

export default Header;
