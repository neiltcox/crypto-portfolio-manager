import React from "react";
import { Route, Switch } from 'react-router-dom';
import { HashRouter } from 'react-router-dom';
import Portfolios from "./Portfolios";

export default function Router(props){
	return (
		<HashRouter basename="/">
			<Switch>
				<Route path='/' component={Portfolios} />
			</Switch>
		</HashRouter>
	);
}