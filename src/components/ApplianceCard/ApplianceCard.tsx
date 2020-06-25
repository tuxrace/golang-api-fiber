import React from "react";
import {
  Box,
  Card,
  CardHeader,
  CardContent,
  makeStyles,
  Typography,
  Grid,
} from "@material-ui/core";
import styles from "./ApplianceCard.styles";

const useStyles = makeStyles(styles);

interface Props {
  data: any;
}

const ApplianceCard: React.FC<Props> = ({ data }) => {
  const appliance = data;
  const classes = useStyles();

  return (
    <Box py={1}>
      <Card className={classes.card}>
        <CardHeader title={`${appliance.brand} ${appliance.model}`} />
        <CardContent>
          <Typography variant="body1">Model: {appliance.model}</Typography>
          <Typography variant="body1">
            Serial Number: {appliance.serialNumber}
          </Typography>
          <Typography variant="body1">Status: {appliance.status}</Typography>
        </CardContent>
      </Card>
    </Box>
  );
};

export default ApplianceCard;
