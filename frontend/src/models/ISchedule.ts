import { DoctorInterface } from "./IDoctor";
import { WorkPlaceInterface } from "./IWorkPlace";

export interface ScheduleInterface {

    ID: number;
    // entity
    DoctorID : number;
    Doctor: DoctorInterface;

    WokrPlaceID : number;
    WorkPlace: WorkPlaceInterface;

    MedActivityID : number;
    MedActivity: WorkPlaceInterface;

    Time: Date;

   //แสดงข้อมูลมาแสดงมาจาก หลังบ้าน
   }