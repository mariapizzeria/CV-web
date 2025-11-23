import { EducationApiResponse, SkillApiResponse, ExperienceApiResponse, SocialApiResponse } from './interfaces.js';

export class Renderer {
    private educationContainer: HTMLElement | null;
    private skillsContainer: HTMLElement | null;
    private experienceContainer: HTMLElement | null;
    private socialContainer: HTMLElement | null;

    constructor() {
        this.educationContainer = document.getElementById('education-container');
        this.skillsContainer = document.getElementById('skills-container');
        this.experienceContainer = document.getElementById('experience-container');
        this.socialContainer = document.getElementById('social-container');
    }

    public createSocialItem(social: SocialApiResponse[]): void {
        if (!this.socialContainer) {
            console.warn('Container not found');
            return;
        }

        let htmlString = '';

        social.forEach((item) => {
            htmlString += `
            <p><a href="${item.link}" target="_blank" rel="noopener noreferrer"><img src="${item.image}" alt="" class="w-9"></a></p>
            `;
        });
        this.socialContainer.innerHTML = htmlString;
    }

    public createExperienceItem(experience: ExperienceApiResponse[]): void {
        if (!this.experienceContainer) {
            console.warn('Container not found');
            return;
        }

        let htmlString = '';

        experience.forEach((item) => {
            htmlString += `
                    <div class="p-4">
                        <h3 class="text-left text-xl mb-6 mt-12">${item.duration}</h3>
                        <h2 class="text-accent text-center text-4xl mb-12">${item.title}</h2>
                        <p class="text-left text-xl mt-4 mb-6">${item.description}</p>
                    </div>
            `;
        });
        this.experienceContainer.innerHTML = htmlString;
    }

    public createSkillItem(skills: SkillApiResponse[]): void {
        if (!this.skillsContainer) {
            console.warn('Skills container not found');
            return;
        }

        let htmlString = '';

        skills.forEach((item) => {
            const skillItems = item.skill.map(skill => `
                <li>&bull; ${skill}</li>
            `).join('');

            htmlString += `
                <div class="border rounded-lg border-black-500 hover:drop-shadow">
                    <div class="p-4">
                        <h2 class="caveat text-accent text-center text-4xl uppercase">${item.type}</h2>
                        <div class="p-6 text-xl">
                            <ul>
                                ${skillItems}
                            </ul>
                        </div>
                    </div>
                </div>
            `;
        });
        this.skillsContainer.innerHTML = htmlString;
    }

    public createEducationItem(edu: EducationApiResponse[]): void {
        if (!this.educationContainer) {
            console.warn('Education container not found');
            return;
        }

        let htmlString = '';

        edu.forEach((item, index) => {
            const isLeft = index % 2 === 0;

            if (isLeft) {
                htmlString += `
                    <div class="timeline-item flex flex-col md:flex-row items-center relative">
                        <div class="timeline-content w-full md:w-1/2 text-right pr-8 md:pr-12 mb-4 md:mb-0">
                            <h3 class="caveat text-2xl md:text-3xl text-accent">${item.duration}</h3>
                            <span class="text-gray-600">${item.college}</span>
                            <p class="font-bold text-xl md:text-2xl text-accent mt-2">${item.course}</p>
                        </div>
                        <div class="absolute left-1/2 transform -translate-x-1/2 flex items-center justify-center w-12 h-12 md:w-16 md:h-16 z-10">
                            <div class="w-4 h-4 md:w-6 md:h-6 bg-accent rounded-full border-4 border-white"></div>
                        </div>
                        <div class="w-full md:w-1/2 pl-8 md:pl-12"></div>
                    </div>
                `;
            } else {
                htmlString += `
                    <div class="timeline-item flex flex-col md:flex-row items-center relative">
                        <div class="w-full md:w-1/2 pr-8 md:pr-12 mb-4 md:mb-0"></div>
                        <div class="absolute left-1/2 transform -translate-x-1/2 flex items-center justify-center w-12 h-12 md:w-16 md:h-16 z-10">
                            <div class="w-4 h-4 md:w-6 md:h-6 bg-accent rounded-full border-4 border-white"></div>
                        </div>
                        <div class="timeline-content w-full md:w-1/2 text-left pl-8 md:pl-12">
                            <h3 class="caveat text-2xl md:text-3xl text-accent">${item.duration}</h3>
                            <span class="text-gray-600">${item.college}</span>
                            <p class="font-bold text-xl md:text-2xl text-accent mt-2">${item.course}</p>
                        </div>
                    </div>
                `;
            }
        });
        this.educationContainer.innerHTML = htmlString;
    }

    public showLoading(containerType: 'education' | 'skills' | 'experience' | 'social'): void {
        let container: HTMLElement | null;

        switch (containerType) {
            case 'education':
                container = this.educationContainer;
                break;
            case 'skills':
                container = this.skillsContainer;
                break;
            case 'experience':
                container = this.experienceContainer;
                break;
            case 'social':
                container = this.socialContainer;
                break;
            default:
                container = null;
        }

        if (!container) {
            console.warn(`${containerType} container not found for loading`);
            return;
        }

        container.innerHTML = `
        <div class="flex justify-center items-center py-8">
            <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-accent"></div>
        </div>`;
    }

    public showError(message: string, containerType: 'education' | 'skills' | 'experience' | 'social'): void {
        let container: HTMLElement | null;

        switch (containerType) {
            case 'education':
                container = this.educationContainer;
                break;
            case 'skills':
                container = this.skillsContainer;
                break;
            case 'experience':
                container = this.experienceContainer;
                break;
            case 'social':
                container = this.socialContainer;
                break;
            default:
                container = null;
        }

        if (!container) {
            console.warn(`${containerType} container not found for error`);
            return;
        }

        container.innerHTML = `
        <div class="flex justify-center items-center py-8">
            <div class="text-red-500 text-center">
                <p class="text-xl">Ошибка загрузки</p>
                <p class="text-sm">${message}</p>
            </div>
        </div>`;
    }
}