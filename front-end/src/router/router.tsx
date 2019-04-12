import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Home from "../views/Home";
import About from "../views/About";
import SignIn from '../views/Sign'
import AppBar from "@material-ui/core/AppBar";

import "./router.scss";

function AppRouter() {
  return (
    <Router>
      <div>
        <AppBar position="static" className="header-navbar">
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/about/">About</Link>
            </li>
            <li>
              <Link to="/signin/">SignIn</Link>
            </li>
          </ul>
        </AppBar>

        <Route path="/" exact component={Home} />
        <Route path="/about/" component={About} />
        <Route path="/signin/" component={SignIn} />
      </div>
    </Router>
  );
}

export default AppRouter;
