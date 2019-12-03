import { Container, Paper } from '@material-ui/core';
import Grid from '@material-ui/core/Grid';
import { withStyles } from '@material-ui/core/styles';
import clsx from 'clsx';
import moment from 'moment';
import PropTypes from 'prop-types';
import React, { Component } from 'react';
import DailySection from '../components/DailySection';
import DataChart from '../components/DataChart';
import TempChart from '../components/TempChart';
import WeatherApi from '../services/weather-api.service';

const styles = theme => ({
  root: {
    display: 'flex',
  },
  appBarSpacer: theme.mixins.toolbar,
  container: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  paper: {
    padding: theme.spacing(2),
    display: 'flex',
    overflow: 'auto',
    flexDirection: 'column',
  },
  fixedHeight: {
    height: 240,
  },
  minHeight: {
    minHeight: 240,
  },
});
export class Dashboard extends Component {
  constructor() {
    super();

    const moments = [
      moment(),
      moment().add(1, 'day'),
      moment().add(2, 'day'),
      moment().add(3, 'day'),
      moment().add(4, 'day'),
    ];
    const days = moments.map(m => m.toISOString());
    this.state = {
      tempSeries: [],
      humiditySeries: [],
      pressureSeries: [],
      rainSeries: [],
      days,
      dailyForecast: null,
    };
  }

  componentDidMount() {
    WeatherApi.getTemperatureSeries().then(series => {
      this.setState({ tempSeries: series });
    });
    WeatherApi.getHumiditySeries().then(series => {
      this.setState({ humiditySeries: series });
    });
    WeatherApi.getPressureSeries().then(series => {
      this.setState({ pressureSeries: series });
    });
    WeatherApi.getRainSeries().then(series => {
      this.setState({ rainSeries: series });
    });
    const { days } = this.state;
    this.getForecast(days[0]);
  }

  getForecast(day) {
    WeatherApi.getDailyForecast(day).then(f => {
      this.setState({ dailyForecast: f });
    });
  }

  render() {
    const { classes } = this.props;

    const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);
    const minHeight = clsx(classes.paper, classes.minHeight);
    const { tempSeries, days, dailyForecast, humiditySeries, pressureSeries, rainSeries } = this.state;

    return (
      <Container maxWidth="lg" className={classes.container}>
        <Grid container spacing={3} className={classes.container}>
          <Grid item xs={12}>
            <Paper className={minHeight}>
              <DailySection days={days} getForecast={this.getForecast.bind(this)} forecast={dailyForecast} />
            </Paper>
          </Grid>

          <Grid item xs={12}>
            <Paper className={fixedHeightPaper}>
              <TempChart series={tempSeries} />
            </Paper>
          </Grid>

          <Grid item xs={4}>
            <Paper className={fixedHeightPaper}>
              <DataChart series={humiditySeries} title="Humidity" yAxisLabel="Percentage %" />
            </Paper>
          </Grid>

          <Grid item xs={4}>
            <Paper className={fixedHeightPaper}>
              <DataChart series={pressureSeries} title="Atmospheric pressure" yAxisLabel="Hectopascals" />
            </Paper>
          </Grid>

          <Grid item xs={4}>
            <Paper className={fixedHeightPaper}>
              <DataChart series={rainSeries} title="Rain volume per 3 hours" yAxisLabel="mm" />
            </Paper>
          </Grid>
        </Grid>
      </Container>
    );
  }
}

Dashboard.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Dashboard);
