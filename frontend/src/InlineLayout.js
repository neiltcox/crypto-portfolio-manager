import React from "react";
import coerceChildArray from "./coerceChildArray";

export default function InlineLayout(props){
	let classes = ['inline_layout'];

	if(['center'].includes(props.align)){
		classes.push('align_'+props.align);
	}

	if(['shim', 'minor', 'major'].includes(props.gap_size)){
		classes.push('gap_size_'+props.gap_size);
	}

	return (
		<div className={classes.join(' ')}>
			{
				coerceChildArray(props.children).map(
					(child, index, child_arr) => {
						let item_classes = ['item'];

						if(index == 0){
							item_classes.push('first');
						}

						return (
							<div
								className={item_classes.join(' ')}
								key={
									(props.keys && props.keys.length == child_arr.length &&
										props.keys[index]
									) || (
										'i_'+index
									)
								}
							>
								{child}
							</div>
						);
					}
				)
			}
		</div>
	);
}