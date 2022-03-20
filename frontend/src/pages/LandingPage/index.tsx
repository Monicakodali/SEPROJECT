import { Container, Box, ContainerProps, Typography, Grid, Button as MuiButton, Divider, Paper, List, ListItem, Link, ListItemText, FormControl, InputLabel, Input, InputAdornment } from '@mui/material';
import { styled } from '@mui/material/styles';
import SearchIcon from '@mui/icons-material/Search';
import React from 'react';
import { useNavigate } from "react-router-dom";

type HeaderProps = {
  maxWidth: ContainerProps['maxWidth'],
  handleSearch: (searchTerm: string) => void
}

const StyledLink = styled(Link)(({ theme }) => {
  return {
    '&:not(:last-child)': {
      marginRight: theme.spacing(3)
    }
  }
})

const Button = styled(MuiButton)(({ theme }) => {
  return {
    marginBottom: theme.spacing(1),
    '&:not(:last-child)': {
      marginRight: theme.spacing(1)
    }
  }
})


const opacity = 0.4

const headerLinks = [
  {label: "Write a Review", href: '#'},
  {label: "Events", href: '#'},
  {label: "Talk", href: '#'},
]

function Header({ maxWidth, handleSearch }: HeaderProps) {

  const [search, setSearch] = React.useState('')

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    handleSearch(search)
  }

  return(
    <Box sx={{background: `url("/images/campus.jpg") rgba(0, 0, 0, ${opacity})`, backgroundBlendMode: 'multiply', backgroundPosition: 'center top', backgroundSize: 'cover', height: 600}}>
    <Container maxWidth={maxWidth} sx={{p: 3, display: 'flex', flexDirection: 'column', height: '100%'}}>
      <Box sx={{display: 'flex', justifyContent: 'space-between'}}>
        <Box style={{paddingTop: 6}}>
          {headerLinks.map(({href, label}) => <StyledLink underline="hover" color="#FFF" href={href}>{label}</StyledLink>)}
        </Box>
        <Box sx={{color: "white"}}>
          <Button color="inherit">
            Log In
          </Button>
          <Button color="inherit" variant="outlined">
            Sign Up
          </Button>
        </Box>
      </Box>
      <Box sx={{flexGrow: 1, display: 'flex', justifyContent: 'center', alignItems: 'center'}}>
        <Box sx={{maxWidth: 700, flexGrow: 1, display: 'flex', flexDirection: 'column', alignItems: 'center'}}>
          <img src="/images/logo.png" alt="logo" style={{width: 250, height:'auto'}} />
            <Paper sx={{width: '100%', display: 'flex', mt: 4}} component="form" onSubmit={handleSubmit}>
              <FormControl fullWidth>
                <Input
                  id="search-input"
                  value={search}
                  onChange={e => setSearch(e.target.value)}
                  placeholder="restaurants, buildings, libraries..."
                  inputProps={{
                    "aria-label": "Search Input"
                  }}
                  startAdornment={<InputAdornment position="start">Find</InputAdornment>}
                  sx={{py: 0.5, px: 1}}
                />
              </FormControl>
              <MuiButton variant="contained" sx={{borderTopLeftRadius: 0, borderBottomLeftRadius: 0}} type="submit">
                <SearchIcon />
              </MuiButton>
            </Paper>
        </Box>
      </Box>
    </Container>
  </Box>
  )
}



export default function LandingPage() {

  const navigate = useNavigate()
 
  return (
    <div>
      <Header handleSearch={s => navigate('/search?query=' + s)} maxWidth="md"/>
      <Container maxWidth="md">
        <Grid container spacing={3}>
          <Grid item xs={4}>
          </Grid>
          <Grid item xs={4}>
          </Grid>
          <Grid item xs={4}>
          </Grid>
        </Grid>
      </Container>
    </div>
  );
};