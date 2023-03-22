import { ActivatedRouteSnapshot, DetachedRouteHandle, RouteReuseStrategy } from '@angular/router'

export class RouteReuseStrategyService implements RouteReuseStrategy {
  /*  execution：
      shouldReuseRoute -> retrieve -> shouldDetach -> store -> shouldAttach -
    -> retrieve(若shouldAttach返回true) -> store(若shouldAttach返回true) 
  */
  public static handlers: { [key: string]: DetachedRouteHandle } = {}
  shouldReuseRoute(future: ActivatedRouteSnapshot, curr: ActivatedRouteSnapshot): boolean {
    return future.routeConfig === curr.routeConfig && JSON.stringify(future.params) === JSON.stringify(curr.params)
  }
  retrieve(route: ActivatedRouteSnapshot): DetachedRouteHandle {
    if (!RouteReuseStrategyService.handlers[this.getRouteUrl(route)]) {
      return null;
    }
    return RouteReuseStrategyService.handlers[this.getRouteUrl(route)];
  }
  shouldDetach(route: ActivatedRouteSnapshot): boolean {
    if (route.data && route.data.keep) {
      return true
    } else {
      return false
    }
  }
  shouldAttach(route: ActivatedRouteSnapshot): boolean {
    return !!RouteReuseStrategyService.handlers[this.getRouteUrl(route)];
  }
  store(route: ActivatedRouteSnapshot, handle: DetachedRouteHandle): void {
    RouteReuseStrategyService.handlers[this.getRouteUrl(route)] = handle;
  }
  /** 使用route的path作为快照的key */
  getRouteUrl(route: ActivatedRouteSnapshot) {
    const path = route['_routerState'].url.replace(/\//g, '_');
    return path;
  }
}