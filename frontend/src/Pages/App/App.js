import React from "react";
import { BrowserRouter } from "react-router-dom";
import "./App.css";
import { AppRoutes } from "../../Routes";
import { HostnamesProvider } from "../../Contexts/Hostnames";

function App() {
  return (
    <HostnamesProvider >
      <BrowserRouter>
        <AppRoutes />
      </BrowserRouter>
    </HostnamesProvider>
  );
}

export default App;
