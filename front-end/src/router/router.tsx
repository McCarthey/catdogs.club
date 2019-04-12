import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Home from '../views/Home'
import About from '../views/About'

import './router.scss'

function AppRouter() {
  return (
    <Router>
      <div>
        <nav className="header-navbar">
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/about/">About</Link>
            </li>
          </ul>
        </nav>

        <Route path="/" exact component={Home} />
        <Route path="/about/" component={About} />
      </div>
    </Router>
  );
}

export default AppRouter