import {render} from "solid-js/web";

import "./index.css";
import App from "./app";
import {lazy} from "solid-js";
import {Router} from "solid-app-router";

const routes = [
    {
        path: "/",
        component: lazy(() => import("./pages/home"))
    },
    {
        path: "/about",
        component: lazy(() => import("./pages/about"))
    },
];

render(() =>
    <Router routes={routes}>
        <App/>
    </Router>, document.getElementById("root")
);
