import React from "react";
import InlineLayout from "./InlineLayout";
import Stack from './Stack';
import ContentBox from "./ContentBox";
import exchangeName from './exchangeName';
import StyledText from "./StyledText";

export default function PortfolioTile(props) {
	return (
		<div className='portfolio_tile'>
			<ContentBox link={"/portfolio/"+props.portfolio.id}>
				<Stack gap_size='minor'>
					<StyledText styling='title'>{props.portfolio.name}</StyledText>
					<InlineLayout gap_size='minor'>
						<StyledText styling='standard'>{exchangeName[props.portfolio.exchange]}</StyledText>
						<StyledText styling='standard'>Connected</StyledText>
					</InlineLayout>
				</Stack>
			</ContentBox>
		</div>
	);
}