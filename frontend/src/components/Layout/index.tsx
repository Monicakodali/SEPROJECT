import * as React from 'react';
import Header from './Header'

import Toolbar from '@mui/material/Toolbar';
import { styled } from '@mui/material/styles';


type LayoutProps = {
  children: React.ReactNode 
}

export default function Layout({children}: LayoutProps) {
  
  return (
    <>
    <Header loggedIn={false}/>
    <Toolbar />
    <main>
      {children}
    </main>
    </>
  );
};