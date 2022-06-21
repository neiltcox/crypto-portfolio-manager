import React from 'react';

export default function FixedWidth(props) {

	let classes = ['fixed_width'];

	if(['small', 'medium', 'large'].includes(props.width)){
		classes.push(props.width+'_width');
	}

	return (
		<div className={classes.join(' ')}>
			{props.children}
		</div>
	);
}