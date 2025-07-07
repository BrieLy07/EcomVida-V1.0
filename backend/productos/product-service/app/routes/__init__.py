from fastapi import APIRouter
from app.controllers.get_all_products import router as get_all
from app.controllers.get_product_by_id import router as get_one
from app.controllers.create_product import router as create
from app.controllers.update_product import router as update
from app.controllers.delete_product import router as delete

router = APIRouter()
router.include_router(get_all)
router.include_router(get_one)
router.include_router(create)
router.include_router(update)
router.include_router(delete)
