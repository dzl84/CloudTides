export const cloudPlatform = {
  ThinkCloud: 'https://pod.thinkcloud.lenovo.com/api',
  Local: 'http://127.0.0.1:3000/api',
};

export const defaultCloudPlatformURL = cloudPlatform.ThinkCloud;

export const resTypes = {
  Fixed: 'Fixed',
  Dynamic: 'Dynamic',
};

export const roleTypes = {
  SITE_ADMIN: 'Site Admin',
  ORG_ADMIN: 'Org Admin',
  USER: 'User',
};

export const roleTypes4Org = {
  ORG_ADMIN: 'Org Admin',
  USER: 'User',
};

export const defaultResType = resTypes.Fixed;
export const defaultRoleType4Site = 'ORG_ADMIN';
export const defaultRoleType4Org = 'USER';