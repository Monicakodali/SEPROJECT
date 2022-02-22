import { Container, Box, ContainerProps, Typography, Grid, Button as MuiButton, Divider, Paper, List, ListItem, Link, ListItemText } from '@mui/material';
import { styled } from '@mui/material/styles';
import axios from 'axios';
import * as React from 'react';
import { useParams } from 'react-router-dom';
import {
  Navigate,
} from "react-router-dom";
import { Rating } from '../../components';
import Tags from '../../components/Tags/index';
import StarOutlineIcon from '@mui/icons-material/StarOutline';
import IosShareIcon from '@mui/icons-material/IosShare';
import BookmarkBorderIcon from '@mui/icons-material/BookmarkBorder';
import CameraAltIcon from '@mui/icons-material/CameraAlt';

type HeaderProps = {
  name: string,
  maxWidth: ContainerProps['maxWidth'],
  rating: number,
  numRatings: number,
  tags: string[]
}

function Header({name, maxWidth, rating, numRatings, tags}: HeaderProps) {
  return(
    <Box sx={{backgroundColor: 'tomato', display: 'flex', alignItems: 'flex-end', height: 300}}>
    <Container maxWidth={maxWidth} sx={{p: 3}}>
      <Typography variant="h1" sx={{color: 'white'}}>{name}</Typography>
      <Box sx={{mt:1, mb: 2}}>
        <Rating size={24} rating={rating} numRatings={numRatings} label={'reviews'} labelStyles={{color: 'white'}} />
      </Box>
      <Box sx={{mb: 2}}>
       <Tags variant="links" tags={tags} linkStyle={{color: 'white'}}/>
      </Box>
      <Box>
        <Typography>More info...</Typography>
      </Box>
    </Container>
  </Box>
  )
}

const Button = styled(MuiButton)(({ theme }) => {
  return {
    marginBottom: theme.spacing(1),
    '&:not(:last-child)': {
      marginRight: theme.spacing(1)
    }
  }
})


export default function Establishment() {
  const { id } = useParams()
  const [data, setData] = React.useState(null)
  const [err, setErr] = React.useState(false)


  // @TODO: populate with actual rating
  // randomly generate rating for now
  const rating = React.useMemo(() => Math.floor(Math.random() * 5) + 1, [])
  // @TODO: populate with actual num ratings
  // randomly generate num of ratings for now
  const numRatings = React.useMemo(() => Math.floor(Math.random() * 100) + 10, [])
  // @TODO: populate with actual isOpen
  // randomly generate whether open or not
  const isOpen = React.useMemo(() => Math.random() > 0.5, [])
  // @TODO: populate tags from DB
  const tags = ['Coffee & Tea', 'Fast Food', 'Example Tag 3']

  React.useEffect(() => {
    if(id) {
      axios.get(`/api/establishments/${id}`).then(res => {
        setData(res.data)
      }).catch(err => {
        setErr(true)
      })
    }
  }, [id])


  if(!id || err) {
    return (<Navigate to="/search" />)
  }

  if(!data) {
    return <div>Loading...</div>
  }


  const { name } = data

  return (
    <Container maxWidth="lg" disableGutters>
      <Header maxWidth="md" {...{name, rating, numRatings, tags}} />
      <Container maxWidth="md">
        <Grid container spacing={3}>
          <Grid item xs={12}>

          </Grid>
          <Grid item xs={8}>
            <Box sx={{mt: 2, mb: 1}}>
              <Button size="small" color="primary" variant="contained" startIcon={<StarOutlineIcon />}>
                Write a Review
              </Button>
              <Button size="small" variant="outlined" startIcon={<CameraAltIcon />}>
                Add Photo
              </Button>
              <Button size="small" variant="outlined" startIcon={<IosShareIcon />}>
                Share
              </Button>
              <Button size="small" variant="outlined" startIcon={<BookmarkBorderIcon />}>
                Save
              </Button>
            </Box>
            <Divider />


          </Grid>
          <Grid item xs={4} >
            <Paper square variant="outlined" sx={{p: 1.5}}>
              <List dense disablePadding>
                <ListItem divider>
                  <ListItemText primary={<Link href={'http://www.example.com'}>http://www.example.com</Link>} />
                </ListItem>
                <ListItem divider>
                  <ListItemText primary={'(352) 371-2323'} />
                </ListItem>
                <ListItem>
                  <ListItemText primary={<Link href={'#'}>Get Directions</Link>} secondary={'123 Main St.'}/>
                </ListItem>
              </List>
            </Paper>
          </Grid>
        </Grid>
      </Container>
    </Container>
  );
};