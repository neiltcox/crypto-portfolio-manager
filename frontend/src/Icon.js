import React from 'react';

export default function Icon(props){
	return (
		<div
			className={
				[
					'icon',
					'icofont-'+props.name,
				].join(' ')
			}
		> </div>
	);
}