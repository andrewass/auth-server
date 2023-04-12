import React from "react";
import {NavigationBarDefault} from "./NavigationBarDefault";
import {NavigationBarMobile} from "./NavigationBarMobile";

export const NavigationBar = () => {

    return (
        <React.Fragment>
            <NavigationBarDefault/>
            <NavigationBarMobile/>
        </React.Fragment>
    );
}