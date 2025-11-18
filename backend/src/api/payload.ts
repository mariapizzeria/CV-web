export interface Education {
    id: number;
    created: Date;
    updated: Date;
    deleted: Date;
    duration: string;
    college: string;
    course: string;
}

export interface EducationApiResponse {
    duration: string;
    college: string;
    course: string;
}