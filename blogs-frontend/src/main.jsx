import React from "react";
import ReactDom from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import App from "./App";
import "./index.css";

ReactDome.CreateRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <BrowserRouter>
    <App/>
    </BrowserRouter>
    </React.StrictMode>
)