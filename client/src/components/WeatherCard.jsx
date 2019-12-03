import { Container } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import React from 'react';

const useStyles = makeStyles(theme => ({
  card: {
    display: 'flex',
    width: '100%',
    height: '100%',
  },
  content: {
    flex: '1 0 ',
  },
  cover: {
    flex: '1 1 ',
    height: '100%',
  },
}));

export default function WeatherCard(props) {
  const classes = useStyles();

  const { description, condition, code } = props;

  return (
    <Container className={classes.card}>
      <div className={classes.content}>
        <Typography component="h5" variant="h5">
          {condition}
        </Typography>
        <Typography variant="subtitle1" color="textSecondary">
          {description}
        </Typography>
      </div>

      <div
        className={classes.cover}
        style={{
          backgroundImage: `url(http://openweathermap.org/img/wn/${code}@2x.png)`,
          backgroundPosition: 'center',
          backgroundSize: 'contain',
          backgroundRepeat: 'no-repeat',
        }}
      />
    </Container>
  );
}
