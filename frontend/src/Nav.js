import React from "react";
import InlineLayout from './InlineLayout';

export default function Nav(props) {
	return (
		<div className='nav'>
			<InlineLayout>
				<div className='logo'>Coinbake</div>
			</InlineLayout>
		</div>
	);
}