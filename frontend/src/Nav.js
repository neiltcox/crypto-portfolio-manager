import React from "react";
import { NavLink } from "react-router-dom";
import InlineLayout from './InlineLayout';
import Icon from './Icon';

export default function Nav(props) {
	return (
		<div className='nav'>
			<InlineLayout>
				<div className='logo'>Coinbake</div>
				<div className='links'>
					<NavLink to="/portfolios" activeClassName="active"><Icon name="chart-pie"/>Portfolios</NavLink>
					<NavLink to="/activity" activeClassName="active"><Icon name="flag"/>Activity</NavLink>
				</div>
			</InlineLayout>
		</div>
	);
}