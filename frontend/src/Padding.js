import React from "react";

export default function Padding(props) {
	let classes = ['padding'];

	if (['minor', 'major'].includes(props.pad_size)) {
		classes.push(props.pad_size+'_pad_size');
	}else{
		classes.push('major_pad_size');
	}

	return (
		<div className={classes.join(' ')}>
			{props.children}
		</div>
	);
}