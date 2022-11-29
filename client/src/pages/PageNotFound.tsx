import { Link } from "react-router-dom";

function PageNotFound() {
  return (
    <div>
      <p>Nothing was found here</p>
      <p>
        <Link to="/">Want to go back?</Link>
      </p>
    </div>
  );
}

export default PageNotFound;
