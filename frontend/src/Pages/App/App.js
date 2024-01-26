import React from "react";
import { BrowserRouter } from "react-router-dom";
import "./App.css";
import { AppRoutes } from "../../Routes";

function App() {
  return (
    <BrowserRouter>
      <AppRoutes />
    </BrowserRouter>
  );
}

export default App;
