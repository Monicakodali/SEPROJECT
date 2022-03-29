import { useState } from 'react';
import { AppBar, Toolbar, Container, Box, Grid, Button, Typography, TextField } from '@mui/material';
import { Link } from '../../components'
import { useNavigate } from 'react-router-dom';
import React from 'react';
import {Route,Navigate} from "react-router-dom";


type HeaderProps = {
  onLogoClick: () => void
}

function Header({onLogoClick}: HeaderProps) {
  return(
    <AppBar position="static">
      <Toolbar style={{minHeight: 64, alignItems: 'center', justifyContent: 'center'}}>
          <img src="/images/logo-white.png" style={{height: 36, width: 'auto', cursor: 'pointer'}}  alt="yelp UF" onClick={() => onLogoClick()} />
      </Toolbar>
    </AppBar>
  )
}


export default function LoginPage() {

  const [u, setU] = useState('')
  const [p, setP] = useState('')
  const [n, setN] = useState('')
  const navigate = useNavigate()
  
  const onSignUp = () => {
    console.log({username: u, name:n,  password: p})
  }
  
  return (
    <div>
      <Header onLogoClick={() => navigate('/')} />
        <Container  sx={{py: 4, minHeight: 720, display: 'flex'}} maxWidth="md">
          <Grid container justifyContent="space-between" spacing={4} style={{flex: 1}}>
            <Grid item xs={6} container direction="column" alignItems="center" justifyContent="center">
              <Typography variant="h4" color="primary" sx={{fontSize: 24, mb: 2, mt: 2}}>Sign up to Yelp UF</Typography>
              
                
              <Box sx={{my: 3, maxWidth: 300}}>
                <TextField inputProps={{'aria-label': 'UF Email'}} placeholder="UF Email" margin="dense" size="small" fullWidth value={u} onChange={e => setU(e.target.value)} />
                <TextField inputProps={{'aria-label': 'Name'}} placeholder="Name" margin="dense" size="small" fullWidth value={n} onChange={e => setN(e.target.value)} />
                <TextField inputProps={{'aria-label': 'Password'}} type="password" placeholder="Password" margin="dense" size="small" fullWidth value={p} onChange={e => setP(e.target.value)} />
                
                <Button sx={{mt: 5}} variant="contained" fullWidth onClick={onSignUp}>
                  Sign Up
                </Button>
              </Box>

            </Grid>
            <Grid item xs={6} container direction="column" alignItems="center" justifyContent="center">
              <img src="/images/tower-circle.png" alt="tower" aria-hidden="true" style={{width: '90%', height: 'auto'}} />
            </Grid>
          </Grid>
        </Container>    
    </div>
  );
};