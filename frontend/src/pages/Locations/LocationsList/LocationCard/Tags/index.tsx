import * as React from 'react';
import { Box, Chip } from '@mui/material';
import { styled } from '@mui/material/styles';


type TagsProps = {
  tags: string[]
}


export default function Tags({tags}: TagsProps) {

  const onClick: React.MouseEventHandler<HTMLDivElement> = (e) => {
    //e.stopPropagation()
  }
  
  return (
    <Box sx={{display: 'flex', flexWrap: 'wrap'}}>
      {tags.map((tag, i) => <Chip size="small" sx={{mr: 1}} onClick={onClick} key={i} label={tag} />)}
    </Box>
  );
};