import React from 'react';

/*

Stylings:
- `standard`: Standard text, default.
- `paragraph`: Body text, with the expectation of multiple lines.
- `note`: Small note text, used for helping the user around the software. Not good for literal content.
- `label`: Labels on fields/data, table headers, etc.
- `title`: Title of some sort of data content.
- `heading`: Headings for sections, usually one per section.

*/

/**
 * Styled Text is a standardized way of styling bits of text.
 */
export default function StyledText(props) {
	let classes = ['text'];
	let customStyling = {};

	if(['standard', 'paragraph', 'note', 'label', 'heading', 'title'].includes(props.styling)){
		classes.push(props.styling);
	}else{
		classes.push('standard');
	}

	if(props.inline){
		classes.push('inline');
	}

	if(['center'].includes(props.align)){
		classes.push('align_'+props.align);
	}

	if(['positive', 'negative'].includes(props.sentiment)){
		classes.push('sentiment_'+props.sentiment);
	}
	
	return (
		<div className={classes.join(' ')} style={customStyling}>
			{props.children}
		</div>
	);
}