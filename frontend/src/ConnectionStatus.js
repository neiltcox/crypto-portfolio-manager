import React from "react";
import InlineLayout from "./InlineLayout";
import StyledText from "./StyledText";

export default function ConnectionStatus(props){
	return (
		<InlineLayout gap_size='shim'>
			<div className='connection_status_indicator' data-sentiment={props.sentiment}></div>
			<StyledText sentiment={props.sentiment}>{props.sentiment == 'positive' ? 'Connected' : 'Disconnected'}</StyledText>
		</InlineLayout>
	);
}