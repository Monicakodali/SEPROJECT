import { Container, Box, ContainerProps, Typography, Grid, Button as MuiButton, Divider, Paper, List, ListItem, ListItemText, FormControl, InputLabel, Input, InputAdornment, Card, CardContent, CardMedia } from '@mui/material';
import { styled } from '@mui/material/styles';
import SearchIcon from '@mui/icons-material/Search';
import React from 'react';
import { useNavigate } from "react-router-dom";
import { Link } from '../../components'

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

const bodyCategories = [
  { name: 'restaurants', image: '/images/swamp.jpg', href: '#'},
  { name: 'buildings', image: '/images/engineering-building.png', href: '#'},
  { name: 'libraries', image: '/images/library.jpg', href: '#'}
]

function Header({ maxWidth, handleSearch }: HeaderProps) {

  const [search, setSearch] = React.useState('')

  const navigate = useNavigate()

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    handleSearch(search)
  }

  return(
    <Box sx={{background: `url("/images/campus.jpg") rgba(0, 0, 0, ${opacity})`, backgroundBlendMode: 'multiply', backgroundPosition: 'center top', backgroundSize: 'cover', height: 600}}>
    <Container maxWidth={maxWidth} sx={{p: 3, display: 'flex', flexDirection: 'column', height: '100%'}}>
      <Box sx={{display: 'flex', justifyContent: 'space-between'}}>
        <Box style={{paddingTop: 6}}>
          {headerLinks.map(({href, label}) => <StyledLink to={href} underline="hover" color="#FFF">{label}</StyledLink>)}
        </Box>
        <Box sx={{color: "white"}}>
          <Button onClick={() => navigate('/login')} color="inherit">
            Log In
          </Button>
          <Button onClick={() => navigate('/signup')} color="inherit" variant="outlined">
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
                  placeholder={bodyCategories.map(c => c.name).join(', ') + "..."}
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
      <Box sx={{backgroundColor: '#f5f5f5', py: 4}}>
        <Container maxWidth="md">
          <Typography variant="h4" color="primary" sx={{textAlign: 'center', fontSize: 20, mb: 3}}>Yelp UF</Typography>
          <Grid container justifyContent="space-between" spacing={4}>
            {bodyCategories.map(c => {

              return (<Grid item key={c.name} sx={{flex: 1}}>
                <Card variant="outlined" sx={{cursor: 'pointer', '&:hover': { borderColor: '#aaa', borderWidth: 1 } }} tabIndex={0} onClick={() => navigate('/search?category=' + c.name)}>
                <CardMedia
                  component="img"
                  height="140"
                  image={c.image}
                  alt={`${c.name} example image`}
                />
                <CardContent sx={{ '&&': { py: 1.5 }}}>
                  <Typography style={{textTransform: 'capitalize', textAlign: 'center'}}>{c.name}</Typography>
                </CardContent>
                </Card>
                </Grid>)
            })}


          </Grid>
        </Container>
      </Box>
      <Box minHeight={300}>
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
      </Box>
      
    </div>
  );
};