import React, { useState } from 'react';
import { AppBar, Toolbar, Container, Box, Grid, Button, Typography, TextField, Alert } from '@mui/material';
import { Link } from '../../components'
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../../context/auth';


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

type Form<T> = {
  username: T
  password: T
}

export default function LoginPage() {

  const [form, setForm] = useState<Form<string>>({
    username: '',
    password: ''
  })
  const [touched, setTouched] = useState<Form<boolean>>({
    username: false,
    password: false
  })
  const [isLoading, setLoading] = useState(false)
  const [error, setError] = useState(false)
  const [success, setSuccess] = useState(false)

  const navigate = useNavigate()
  
  const { login, isAuthenticated } = useAuth()

  React.useEffect(() => {
    if(!success) {
      return
    }
    const t = setTimeout(() => {
      navigate('/')
    }, 3000)

    return () => clearTimeout(t)

  }, [success, navigate])

  React.useEffect(() => {
    if(isAuthenticated && !success) {
      navigate('/', { replace: true })
    }
  }, [isAuthenticated, navigate, success])

  


  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm(f => ({...f, [e.target.name]: e.target.value}))
    setTouched(f => ({...f, [e.target.name]: true}))
  }

  const onLogIn = () => {
    setTouched({username: true, password: true})
    setLoading(true)
    setError(false)
    login(form.username, form.password).then((result) => {
      setLoading(false)
      setError(!result)
      setSuccess(result)
    })
  }
  
  return (
    <div>
      <Header onLogoClick={() => navigate('/')} />
        <Container  sx={{py: 4, minHeight: 720, display: 'flex'}} maxWidth="md">
          <Grid container justifyContent="space-between" spacing={4} style={{flex: 1}}>
            <Grid item container alignItems="flex-end" xs={12}>
              {error && <Alert onClose={() => setError(false)} severity="error" sx={{flex: 1}}>The email address or password you entered is incorrect.</Alert>}
              {success && <Alert severity="success" sx={{flex: 1}}>Successfully logged in. Redirecting...</Alert>}
            </Grid>
            <Grid item xs={6} container direction="column" alignItems="center" justifyContent="center">
              <Typography variant="h4" color="primary" sx={{fontSize: 24, mb: 2, mt: 2}}>Log in to Yelp UF</Typography>
              
                <Typography sx={{'& a': { fontWeight: 'bold'}}}>
                  New to Yelp UF? <Link to="/signup">Sign up</Link>
                </Typography>
              <Box component="form" sx={{my: 3, maxWidth: 300}}>
                <TextField name="username" autoComplete="username" spellCheck={false} disabled={isLoading || success} error={!!form.username && touched.username && error} inputProps={{'aria-label': 'UF Email'}} placeholder="UF Email" margin="dense" size="small" fullWidth value={form.username} onChange={onChange} />
                <TextField name="password" autoComplete="current-password" disabled={isLoading || success} error={!!form.password && touched.password && error} inputProps={{'aria-label': 'Password'}} type="password" placeholder="Password" margin="dense" size="small" fullWidth value={form.password} onChange={onChange} />
                <div style={{textAlign: 'right'}}>
                  <Link sx={{fontSize: 12}} to="/forgot-password">Forgot Password?</Link>
                </div>
                <Button sx={{mt: 5}} variant="contained" fullWidth onClick={onLogIn}>
                  Log In
                </Button>
                <div style={{textAlign: 'right'}}>
                  <Typography color="textSecondary" variant="caption">New to Yelp UF? <Link to="/signup">Sign Up</Link></Typography>
                </div>
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