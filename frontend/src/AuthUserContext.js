import React from 'react';

const AuthUserContext = React.createContext(
	{
		user: {},
		authenticated: true,
	}
);

export default AuthUserContext;
