import React from 'react';
import Router from './Router';
import Nav from './Nav';
import FixedWidth from './FixedWidth';

import AuthUserContext from './AuthUserContext';

export default function App(props) {
	return (
		<AuthUserContext.Provider value={{user: {}, authenticated: false}}>
			<FixedWidth width='medium'>
				<Router/>
			</FixedWidth>
		</AuthUserContext.Provider>
	);
}