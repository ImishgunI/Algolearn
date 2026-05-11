import { http } from "./http";

export interface Course {
  id: number;
  title: string;
  description: string;
  // другие поля
}

export async function getAllCourses(): Promise<Course[]> {
  return http("/api/courses");
}