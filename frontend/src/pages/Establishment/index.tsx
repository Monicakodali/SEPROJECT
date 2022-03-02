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
import useDiningHours from '../../hooks/dininghours';
import { addDays, getDay, isAfter, isBefore, isValid, parse } from 'date-fns';
import MiniMap from './MiniMap'
import useBuilding from '../../hooks/buildings/index';

type HeaderProps = {
  name: string,
  maxWidth: ContainerProps['maxWidth'],
  rating: number,
  numRatings: number,
  tags: string[],
  isOpen: boolean | null,
  hoursOfOperation: Date[] | null
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
  const [data, setData] = React.useState<null | Establishment>(null)
  const [err, setErr] = React.useState(false)

  const hours = useDiningHours(data)
  const building = useBuilding(data)

  // @TODO: populate with actual rating
  // randomly generate rating for now
  const rating = React.useMemo(() => Math.floor(Math.random() * 5) + 1, [])
  // @TODO: populate with actual num ratings
  // randomly generate num of ratings for now
  const numRatings = React.useMemo(() => Math.floor(Math.random() * 100) + 10, [])
  // @TODO: populate tags from DB
  const tags = ['Coffee & Tea', 'Fast Food', 'Example Tag 3']


  const day = (getDay(new Date()) + 6) % 7
  
  // for today
  const hoursOfOperation = React.useMemo<null | Date[]>(() => {
    if(!hours) {
      return null
    } else {
      const h = hours[day]
      if(h.hoursOfOperation === 'CLOSED') {
        return null
      } else {
        const [start, end] = h.hoursOfOperation.split(' ‑ ')
        let open = parse(start, 'h:mm aa', new Date())
        let close = parse(end, 'h:mm aa', new Date())
        if(!isValid(open) || !isValid(close)) {
          return null
        }
        if(isBefore(close, open)) {
          close = addDays(close, 1)
        }
        return [open, close]
      }

    }

  }, [hours, day])

  const isOpen = React.useMemo(() => {
    if(!hoursOfOperation) {
      return null
    } else {
      const now = new Date()
      return isAfter(now, hoursOfOperation[0]) && isBefore(now, hoursOfOperation[1])
    }
  }, [hoursOfOperation])
 
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

  console.log({data, hoursOfOperation, isOpen, hours, building})

  const { name } = data

  return (
    <Container maxWidth="lg" disableGutters>
      <Header maxWidth="md" {...{name, rating, numRatings, tags, isOpen, hoursOfOperation}} />
      <Container maxWidth="md">
        <Grid container spacing={3}>
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
            <Box sx={{mt: 2, mb: 1}}>
              <Typography component="h2" variant="h5" gutterBottom>Location &amp; Hours</Typography>
              <Grid container spacing={3}>
                <Grid item xs={5}>
                  {data?.x && data?.y && <MiniMap height={150} coordinates={[data.x, data.y]} />}


                  <Box sx={{display: 'flex', alignItems: 'flex-start', py: 2, justifyContent: 'space-between' }}>
                    {building !== null && <Box sx={{mr: 1}}>
                      <Typography style={{lineHeight: 1}}>{building.PropName}</Typography>
                      <Typography variant="caption">
                        Gainesville, FL
                      </Typography>
                    </Box>}
                    <Button variant="contained" size="small">
                      Directions
                    </Button>
                  </Box>

                </Grid>
                {hours !== null && <Grid item xs={7}>
                  {
                    hours.map((hr, i) => {

                      const fontWeight = i === day ? 'bold' : 'normal'

                      return <Box key={hr.day} sx={{display: 'flex' }}>
                      <Typography sx={{flex: '0 1 60px', fontWeight}}>
                        {hr.day.slice(0, 3)}
                      </Typography>
                      <Typography sx={{flex: '1 1 auto', fontWeight}}>
                        {hr.hoursOfOperation}
                      </Typography>
                      {i === day && <Typography sx={{flex: '1 1 auto', fontWeight: 'bold', color: isOpen ? 'green' : 'red'}}>
                        {isOpen ? 'Open now' : 'Closed'}
                      </Typography>}
                    </Box>
                    })
                  }
                </Grid>}
              </Grid>


            </Box>
            <Divider />
            <Box sx={{mt: 2, mb: 1}}>
              <Typography component="h2" variant="h5">Reviews</Typography>


            </Box>



          </Grid>
          <Grid item xs={4} >
            <Paper square variant="outlined" sx={{p: 1.5, mt: 3}}>
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