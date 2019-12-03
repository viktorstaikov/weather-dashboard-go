import Box from '@material-ui/core/Box';
import { withStyles } from '@material-ui/core/styles';
import Tab from '@material-ui/core/Tab';
import Tabs from '@material-ui/core/Tabs';
import Typography from '@material-ui/core/Typography';
import moment from 'moment';
import PropTypes from 'prop-types';
import React, { Component } from 'react';
import SwipeableViews from 'react-swipeable-views';
import Forecast from './Forecast';
import Loading from './Loading';

function TabPanel(props) {
  const { children, value, index, ...other } = props;

  return (
    <Typography
      component="div"
      role="tabpanel"
      hidden={value !== index}
      id={`full-width-tabpanel-${index}`}
      aria-labelledby={`full-width-tab-${index}`}
      {...other}
    >
      <Box p={3}>{children}</Box>
    </Typography>
  );
}

TabPanel.propTypes = {
  children: PropTypes.node,
  index: PropTypes.any.isRequired,
  value: PropTypes.any.isRequired,
};

function a11yProps(index) {
  return {
    id: `full-width-tab-${index}`,
    key: `full-width-tab-${index}`,
    'aria-controls': `full-width-tabpanel-${index}`,
  };
}

const styles = theme => ({
  root: {
    backgroundColor: theme.palette.background.paper,
  },
  minHeight: {
    minHeight: '240px',
  },
});

export class DailySection extends Component {
  constructor() {
    super();

    this.state = { value: 0 };
  }

  handleChange(event, newValue) {
    const day = this.props.days[newValue];
    this.props.getForecast(day);
    this.setState({ value: newValue });
  }

  handleChangeIndex(index) {
    this.setState({ value: index });
  }

  render() {
    const { classes, days, forecast } = this.props;

    const today = moment();
    const labels = days.map(d => {
      const m = moment(d);
      const day = m.isSame(today, 'day') ? 'Today' : m.format('dddd');
      const date = m.format('DD.MM');
      return (
        <div>
          <Typography variant="h6">{day}</Typography>
          <Typography variant="subtitle2">{date}</Typography>
        </div>
      );
    });
    const tabs = labels.map((l, idx) => <Tab label={l} {...a11yProps(idx)} />);
    const panels = days.map((d, idx) => {
      const content = this.state.value === idx && forecast ? <Forecast data={forecast} /> : <Loading loading={true} />;

      return (
        <TabPanel className={classes.minHeight} value={this.state.value} index={idx} key={idx} dir="ltr">
          {content}
        </TabPanel>
      );
    });

    return (
      <div className={classes.root}>
        <Tabs
          value={this.state.value}
          onChange={this.handleChange.bind(this)}
          indicatorColor="primary"
          textColor="primary"
          variant="fullWidth"
          aria-label="full width tabs example"
        >
          {tabs}
        </Tabs>

        <SwipeableViews axis="x" index={this.state.value} onChangeIndex={this.handleChangeIndex.bind(this)}>
          {panels}
        </SwipeableViews>
      </div>
    );
  }
}

DailySection.propTypes = {
  days: PropTypes.array.isRequired,
  getForecast: PropTypes.func.isRequired,
  forecast: PropTypes.object,
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(DailySection);
