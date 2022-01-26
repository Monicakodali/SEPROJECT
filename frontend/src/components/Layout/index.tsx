import * as React from 'react';
import Header from './Header'

type LayoutProps = {
  children: React.ReactNode 
}

export default function Layout({children}: LayoutProps) {
  
  return (
    <>
    <Header loggedIn={false}/>
    <main>
      {children}
    </main>
    </>
  );
};