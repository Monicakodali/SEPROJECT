import React, { useState } from 'react';
import { Container, Box, Grid, Button, Typography, TextField, Alert } from '@mui/material';
import { SimpleHeader, Link } from '../../components'
import validate from 'validate.js'
import { useAuth } from '../../context/auth';
import { useNavigate } from 'react-router-dom';

type Field = 'username' | 'password' | 'name' | 'confirmPassword'

type Form<T> = Record<Field, T>

var constraints = {
  username: {
    presence: {
      allowEmpty: false
    },
  },
  password: {
    presence: {
      allowEmpty: false
    },
  },
  confirmPassword: {
    presence: {
      allowEmpty: false
    },
    equality: "password"
  },
  name: {
    presence: {
      allowEmpty: false
    },
  },
};


export default function SignUpPage() {

  const [form, setForm] = useState<Form<string>>({
    username: '',
    password: '',
    name: '',
    confirmPassword: ''
  })
  const [touched, setTouched] = useState<Form<boolean>>({
    username: false,
    password: false,
    name: false,
    confirmPassword: false
  })
  const [errors, setErrors] = useState<Form<string[]>>({
    username: [],
    password: [],
    name: [],
    confirmPassword: []
  })



  const [init, setInit] = useState(false)
  const [error, setError] = useState(false)
  const [success, setSuccess] = useState(false)
  const [isLoading, setLoading] = useState(false)

  const { isAuthenticated, signUp } = useAuth()
  const navigate = useNavigate()

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
    if(isAuthenticated && !success && !init) {
      navigate('/', { replace: true })
    }
  }, [isAuthenticated, navigate, success, init])

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm(f => ({...f, [e.target.name]: e.target.value}))
    setTouched(f => ({...f, [e.target.name]: true}))
  }
  
  const onSignUp = () => {
    setInit(true)
    setTouched({username: true, password: true, name: true, confirmPassword: true})
    const result = validate(form, constraints)
    if(result) {
      setErrors(e => ({...e, ...result}))
      return
    }

    setLoading(true)
    setError(false)
    signUp({
      Username: form.username,
      Password: form.password,
      Name: form.name
    }).then((result) => {
      setError(!result)
      setSuccess(result)
    }).catch(() => {
      setError(true)
    }).finally(() => {
      setLoading(false)
    })
  }

  
  const hasError = (field: Field): boolean => {
    if(error) {
      return true
    }
    return !!errors[field] && errors[field].length > 0
  }
  const getHelperText = (field: Field): string|undefined => {
    return errors[field]?.[0]
  }
  
  return (
    <div>
      <SimpleHeader />
        <Container  sx={{py: 4, minHeight: 720, display: 'flex'}} maxWidth="md">
          <Grid container justifyContent="space-between" spacing={4} style={{flex: 1}}>
            <Grid item container alignItems="flex-end" xs={12}>
              {error && <Alert onClose={() => setError(false)} severity="error" sx={{flex: 1}}>There was an error creating your account.</Alert>}
              {success && <Alert severity="success" sx={{flex: 1}}>Your account was created successfully. Redirecting...</Alert>}
            </Grid>
            <Grid item xs={6} container direction="column" alignItems="center" justifyContent="center">
              <Typography variant="h4" color="primary" sx={{fontSize: 24, mb: 2, mt: 2}}>Sign Up for Yelp UF</Typography>      
              <Box sx={{my: 3, maxWidth: 300}}>
                <TextField required name="username" inputProps={{'aria-label': 'Username'}} placeholder="Username" margin="dense" size="small" fullWidth value={form.username} onChange={onChange} disabled={isLoading || success} error={(!!form.username || init) && touched.username && hasError('username')} helperText={getHelperText('username')}/>
                <TextField required name="name" inputProps={{'aria-label': 'Name'}} placeholder="Name" margin="dense" size="small" fullWidth value={form.name} onChange={onChange} disabled={isLoading || success} error={(!!form.name || init) && touched.name && hasError('name')}  helperText={getHelperText('name')} />
                <TextField required name="password" inputProps={{'aria-label': 'Password'}} type="password" placeholder="Password" margin="dense" size="small" fullWidth value={form.password} onChange={onChange} disabled={isLoading || success} error={(!!form.password || init) && touched.password && hasError('password')}  helperText={getHelperText('password')} />
                <TextField required name="confirmPassword" inputProps={{'aria-label': 'Confirm Password'}} type="password" placeholder="Confirm Password" margin="dense" size="small" fullWidth value={form.confirmPassword} onChange={onChange} disabled={isLoading || success} error={(!!form.confirmPassword || init) && touched.confirmPassword && hasError('confirmPassword')}  helperText={getHelperText('confirmPassword')} />
                
                <Button sx={{mt: 5, mb: 1}} variant="contained" fullWidth onClick={onSignUp}>
                  Sign Up
                </Button>
                <div style={{textAlign: 'right'}}>
                  <Typography color="textSecondary" variant="caption">Already on Yelp UF? <Link to="/login">Log in</Link></Typography>
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