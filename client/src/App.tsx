import "./App.css";
import { Route, Routes } from "react-router-dom";
import Header from "./components/header";
import PageNotFound from "./pages/PageNotFound";
import Home from "./pages/Home";

function App() {
  // the header tag can be extracted into a navbar for better layout
  return (
    <div className="App">
      <Header />
      <div className="content">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="*" element={<PageNotFound />} />
        </Routes>
      </div>
    </div>
  );
}

export default App;
