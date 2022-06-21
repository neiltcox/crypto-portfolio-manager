import React from 'react';

/**
 * A vertical stack of cells, where spacing is handled automatically.
 */
export default function Stack(props){

	let classes = ['stack'];

	if(['major', 'minor', 'shim'].includes(props.gap_size)){
		classes.push('gap_size_'+props.gap_size);
	}

	return (
		<div className={classes.join(' ')}>
			{
				(
					(props.children &&
						(Array.isArray(props.children) &&
							// Children is an array, which is the format we want
							props.children
						) || (
							// Children is not an array, coerce into an array
							[props.children]
						)
					) || (
						// Children is null/undefined, use an empty array
						[]
					)
				).filter(
					(child) => {
						return child != null;
					}
				).map(
					(child, index, child_arr) => {
						let stack_item_classes = ['stack_item'];

						if(index == 0){
							stack_item_classes.push('first');
						}

						return (
							<div
								className={stack_item_classes.join(' ')}
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