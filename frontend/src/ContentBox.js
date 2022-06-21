import React from "react";
import { Link } from "react-router-dom";
import Padding from './Padding';

export default function ContentBox(props){
	let contentBox = (
		<div className='content_box'>
			<Padding>
				{props.children}
			</Padding>
		</div>
	);
	
	if (props.link) {
		contentBox = (
			<Link to={props.link}>
				{contentBox}
			</Link>
		);
	}

	return contentBox;
}