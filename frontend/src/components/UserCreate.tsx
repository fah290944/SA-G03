import React, { useEffect } from "react";

import { Link as RouterLink } from "react-router-dom";

import TextField from "@mui/material/TextField";

import Button from "@mui/material/Button";

import FormControl from "@mui/material/FormControl";

import Container from "@mui/material/Container";

import Paper from "@mui/material/Paper";

import Grid from "@mui/material/Grid";

import Box from "@mui/material/Box";

import Typography from "@mui/material/Typography";

import Divider from "@mui/material/Divider";

import Snackbar from "@mui/material/Snackbar";

import MuiAlert, { AlertProps } from "@mui/material/Alert";

import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";

import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";

import { DatePicker } from "@mui/x-date-pickers/DatePicker";
//สี
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { green } from "@mui/material/colors";
//timedate
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DateTimePicker } from "@mui/x-date-pickers/DateTimePicker";
//combobox
import { ComboBoxComponent } from "@syncfusion/ej2-react-dropdowns";

import { WorkPlaceInterface } from "../models/IWorkPlace";

import { MedActivityInterface } from "../models/IMedActivity";
import { Schedule } from "@mui/icons-material";
import { ScheduleInterface } from "../models/ISchedule";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MenuItem from '@mui/material/MenuItem';
import { DoctorInterface } from "../models/IDoctor";


const theme = createTheme({
  palette: {
    primary: {
      // Purple and green play nicely together.
      main: green[500],
    },
    secondary: {
      // This is green.A700 as hex.
      main: "#e8f5e9",
    },
  },
});

//เด้งบันทึกสำเร็จ ไม่สำเร็จ
const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,

  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function UserCreate() {
 

  //  const [date, setDate] = React.useState<Date | null>(null);

  const [schedule, setSchedule] = React.useState<Partial<ScheduleInterface>>({
    WokrPlaceID: 0,
    MedActivityID: 0,
  });


  //Partial คือเลือกค่า set ค่าได้เฉพาะตัวได้

  const [success, setSuccess] = React.useState(false);

  const [error, setError] = React.useState(false);

  const [medactivity, setmedactivity] = React.useState<MedActivityInterface[]>(
    []
  );
  //เราส่งมาในรูปแบบอาเรย์ ทำการดึงข้อมูล
  const [workplace, setworkplace] = React.useState<WorkPlaceInterface[]>([]);
  const [selectedDate, setSelectedDate] = React.useState<Date | null>(
    new Date()
  );
  const [user, setUser] = React.useState<DoctorInterface>();

  const handleDateChange = (date: Date | null) => {
    setSelectedDate(date);
  };

  const getWorkPlace = async () => {
    const apiUrl = `http://localhost:8080/workPlace`;

    const requestOptions = {
      method: "GET",

      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };
    //การกระทำ //json
    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        console.log(res.data); //show ข้อมูล

        if (res.data) {
          setworkplace(res.data);
        } else {
          console.log("else");
        }
      });
  };
  //activity
  const getMedactivity = async () => {
    const apiUrl = `http://localhost:8080/medActivitys`;

    const requestOptions = {
      method: "GET",

      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };
    //การกระทำ
    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        console.log(res.data);

        if (res.data) {
          setmedactivity(res.data);
        } else {
          console.log("else");
        }
      });
  };
  //activity

  //เปิดปิดตัว Alert
  const handleClose = (
    event?: React.SyntheticEvent | Event,

    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }

    setSuccess(false);
    setError(false);
  };

  console.log(schedule);

  //ทุกครั้งที่พิมพ์จะทำงานเป็น state เหมาะสำหรับกับคีย์ textfield
  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof UserCreate;

    const { value } = event.target;

    //  setUser({ ...user, [id]: value });
  };
  
//กดเลือกคอมโบไม่ได้
  const handleChange = (
    event: SelectChangeEvent<number>
  ) => {
    const name = event.target.name as keyof typeof schedule;
    setSchedule({
      ...schedule,
      [name]: event.target.value,
    });
  };

  // const handleChange = (event: SelectChangeEvent<{ name?: string; value: unknown }>) => {
  //   const name = event.target.name as keyof typeof schedule;
  //     setSchedule({
  //       ...schedule,
  //       [name]: event.target.value,
  //     });
  // };

  function submit() {
    let data = {
      //กับ คอน บรรทัด 48-50 แค่ข้างหน้า ชื่อต้องตรง!!!!!!!
      DoctorID: user?.ID,

      WorkPlaceID: schedule.WokrPlaceID,

      MedActivityID: schedule.MedActivityID,

      Time: selectedDate,
    };

    console.log(data)
    const apiUrl = "http://localhost:8080/saveschedule";

    const requestOptions = {
      method: "POST",

      headers: 
      {  Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json" 
      },

      body: JSON.stringify(data),
      //แปลงข้อมูล
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        if (res.data) {
          setSuccess(true);
        } else {
          setError(true);
        }
      });
  }

  useEffect(() => {
    const getToken = localStorage.getItem("token");
    if (getToken) {
        setUser(JSON.parse(localStorage.getItem("user") || ""));
    }
    //เอามาใช้ๆๆๆๆๆๆๆ*********
    getWorkPlace();
    getMedactivity();
  }, []);

  return (
    <ThemeProvider theme={theme}>
      <Container maxWidth="md">
        <Snackbar
          open={success}
          autoHideDuration={6000}
          onClose={handleClose}
          anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
        >
          <Alert onClose={handleClose} severity="success">
            บันทึกข้อมูลสำเร็จ
          </Alert>
        </Snackbar>

        <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
          <Alert onClose={handleClose} severity="error">
            บันทึกข้อมูลไม่สำเร็จ
          </Alert>
        </Snackbar>

        <Paper>
          <Box
            display="flex"
            sx={{
              marginTop: 2,
            }}
          >
            <Box sx={{ paddingX: 2, paddingY: 1 }}>
              <Typography
                component="h2"
                variant="h6"
                color="inherit"
                gutterBottom
              >
                ตารางเวลาแพทย์
              </Typography>
            </Box>
          </Box>

          <Divider />

          <Grid container spacing={3} sx={{ padding: 2 }}>
            <Grid item xs={6}>
              <p>สถานที่การทำงาน</p>

              <FormControl fullWidth variant="outlined">
                <Select
                  value = {schedule.WokrPlaceID}
                  onChange = {handleChange}
                  inputProps={{
                    name: "WokrPlaceID",
                  }}
                  // defaultValue={0}
                >
                  <MenuItem value={0} key={0}>
                    เลือกสถานที่
                  </MenuItem>
                  {workplace.map((item: WorkPlaceInterface) => (
                    <MenuItem value={item.ID}>{item.Name}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={6}></Grid>

            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined">
                <p>กิจกรรมของแพทย์</p>

                <Select
                  value={schedule.MedActivityID}
                  onChange={handleChange}
                  inputProps={{
                    name: "MedActivityID",
                  }}
                  // defaultValue={0}
                >
                  <MenuItem value={0} key={0}>
                    เลือกกิจกรรม
                  </MenuItem>
                  {medactivity.map((item: MedActivityInterface) => (
                    <MenuItem value={item.ID}>{item.Name}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>
            {/* //วันที่และเวลา */}
            <Grid item xs={7}>
              <FormControl fullWidth variant="outlined">
                <p>วันที่และเวลา</p>

                <LocalizationProvider dateAdapter={AdapterDayjs}>
                  <DateTimePicker
                    renderInput={(props) => <TextField {...props} />}
                    label="กรุณาเลือกวันและเวลา"
                    value={selectedDate} //แก้
                    // onChange={(newValue) => {
                    // setDate(newValue);

                    // }}
                    onChange={setSelectedDate}
                  />
                </LocalizationProvider>
              </FormControl>
            </Grid>

            <Grid item xs={12}>
              <Button component={RouterLink} to="/" variant="contained">
                <Typography
                  color="secondary"
                  component="div"
                  sx={{ flexGrow: 1 }}
                >
                  ย้อนกลับ
                </Typography>
              </Button>

              <Button
                style={{ float: "right" }}
                onClick={submit}
                variant="contained"
                color="primary"
              >
                <Typography
                  color="secondary"
                  component="div"
                  sx={{ flexGrow: 1 }}
                >
                  บันทึกตารางเวลาแพทย์
                </Typography>
              </Button>
            </Grid>
          </Grid>
        </Paper>
      </Container>
    </ThemeProvider>
  );
}

export default UserCreate;
