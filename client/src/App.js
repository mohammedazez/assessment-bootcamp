import React from "react";
import "./App.css";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";


// pages
import Dashboard from "../src/pages/dashboard.jsx"


function App() {
  return (
    <div>
      <Router>
          <Switch>
              <Route path="/">
                  <Dashboard/>
              </Route>
          </Switch>
      </Router>
    </div>
  );
}

export default App;
