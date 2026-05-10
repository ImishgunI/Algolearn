import { http } from "./http";

export interface Lesson {
  id: number;
  course_id: number;
  title: string;
  theory: string;
  algorithm_type: string;
  order_index: number;
}

export async function getLessonsByCourse(courseID: number) {
  return http(`/api/courses/${courseID}/lessons`);
}

export async function getLesson(id: number): Promise<Lesson> {
  return http(`/api/lessons/${id}`);
}

export async function getComments(lessonID: number) {
  return http(`/api/lessons/${lessonID}/comments`);
}

export async function createComment(lessonID: number, body: string) {
  return http(`/api/lessons/${lessonID}/comments`, {
    method: "POST",
    body: JSON.stringify({ body }),
  });
}

export async function getFavoriteStatus(lessonID: number) {
  return http(`/api/lessons/${lessonID}/favorite`);
}

export async function toggleFavorite(lessonID: number) {
  return http(`/api/lessons/${lessonID}/favorite`, { method: "POST" });
}